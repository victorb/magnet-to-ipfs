package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"code.cloudfoundry.org/bytefmt"
	"github.com/anacrolix/missinggo"
	"github.com/anacrolix/torrent"
	api "github.com/ipfs/go-ipfs-api"
)

// Should end up with
// => ipfs add -r -w Sintel
// added QmNXQpEuxbxwQXxVLZvTRrUiomqqwbXbGSva8SSHPWh9SU Sintel/Sintel.de.srt
// added QmXExs5oDeQwFDN6qRbrrx2JwP2Mho2P4SdoZZNsJyPrF1 Sintel/Sintel.en.srt
// added QmPsiMQSVruocq3TG1SmLz2phsBa4XxMadU5f1xyMofHNp Sintel/Sintel.es.srt
// added QmdeXx3KWN7LCEECCrT6uHeH6tuvH6eVJRhj5kaDU8LHLS Sintel/Sintel.fr.srt
// added QmVovwU6cZXRNCsse4gGibrienX6ZcxWVC9W7FTYeJ2KHh Sintel/Sintel.it.srt
// added QmcPZGXSyaQSnDVEtJrgtp2EyxhENiBdpGa8pfs5rMBCWr Sintel/Sintel.mp4
// added QmfDZuTjynaCahG7yZrMHBfTccpCHXgs7MHRhFgieLNK5N Sintel/Sintel.nl.srt
// added QmPwfmpPuqftfNVnfL5UYVCE5eegSk4jpzxgEtGHnR4oMv Sintel/Sintel.pl.srt
// added QmWbquQKiLTPqibijrEeAv8RhRBqDPdehWwNHE3GMVStuH Sintel/Sintel.pt.srt
// added QmVDx5LAHNrdej3HXYSqjEX1SYyEtVQ1khqiJ7D8nCG3vP Sintel/Sintel.ru.srt
// added QmX5YC15NYxJBzWCWTY6c9PusSefJCTNNZg7VfpndT32aH Sintel/poster.jpg
// added QmXgfYJiG3JttNqVHERfGysc1semSrFnhtAiFU1C6oLNWW Sintel
// added QmYSfiNmznRYnSeUQiLVxqzuRwvrJXJLp5ESNb4u5WL1U2

type File struct {
	Path string
	Hash string
}

func main() {
	ipfsClient := api.NewShell("localhost:5001")

	if len(os.Args[1:]) < 1 {
		log.Fatal("Need a Magnet URL as first argument")
	}
	magnetLink := os.Args[1:][0]

	// Initialize torrent client
	// TODO should be in memory-store
	// clientConfig := torrent.Config{DefaultStorage: storage.NewBoltDB(os.TempDir())}
	client, err := torrent.NewClient(nil) // Use default client config
	if err != nil {
		panic(err)
	}
	torrentToDownload, err := client.AddMagnet(magnetLink)
	// torrentToDownload, err := client.AddTorrentFromFile("./KnightsOfTheOldRepublic_dark_334_archive.torrent")
	fmt.Println("Fetching info...")
	// Wait for getting torrent metadata
	<-torrentToDownload.GotInfo()

	// Print how much is left to download
	go func() {
		for {
			missing := uint64(torrentToDownload.BytesMissing())
			missingStr := strconv.FormatUint(missing, 10)
			fmt.Println(bytefmt.ByteSize(missing) + " (" + missingStr + "B) left")
			time.Sleep(time.Second * 10)
		}
	}()
	// Start downloading all files
	torrentToDownload.DownloadAll()

	r := torrentToDownload.NewReader()
	defer func() {
		err := r.Close()
		if err != nil {
			panic(err)
		}
	}()
	// Create empty directory to append all files to
	dirHash, err := ipfsClient.NewObject("unixfs-dir")
	for _, file := range torrentToDownload.Files() {
		// For each file, create reader that we can use for adding to IPFS
		reader := file.NewReader()
		defer func() {
			err := reader.Close()
			if err != nil {
				panic(err)
			}
		}()
		// get reader of just the part that we're reading
		hash, err := ipfsClient.Add(missinggo.NewSectionReadSeeker(r, file.Offset(), file.Length()))
		if err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("added", hash, file.Path())
		// Add to empty directory
		dirHash, err = ipfsClient.PatchLink(dirHash, file.Path(), hash, true)
	}
	fmt.Println("added", dirHash)
	// Wait for all files to have finished
	client.WaitAll()
}
