## magnet-to-ipfs

> Takes a magnet link, downloads the parts and throws it into IPFS

### Requirements

* IPFS running in the background

### Installation

- `go get -u -v github.com/victorbjelkholm/magnet-to-ipfs`

Or download the binaries directly from the Releases page.

### Usage

- Run IPFS in a terminal tab
  - `ipfs daemon --init`
- Run `magnet-to-ipfs` with your magnet link

Example with downloading Sintel (a free, Creative Commons movie):

```
$ magnet-to-ipfs "magnet:?xt=urn:btih:08ada5a7a6183aae1e09d831df6748d566095a10&dn=Sintel&tr=udp%3A%2F%2Fexplodie.org%3A6969&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Ftracker.empire-js.us%3A1337&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337&tr=wss%3A%2F%2Ftracker.btorrent.xyz&tr=wss%3A%2F%2Ftracker.fastcast.nz&tr=wss%3A%2F%2Ftracker.openwebtorrent.com&ws=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2F&xs=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2Fsintel.torrent"

Fetching info...
2018/01/18 14:40:09 error announcing "Sintel" to DHT: no initial nodes
2018/01/18 14:40:09 error bootstrapping dht: no initial nodes
123.3M (129302391B) left
added QmNXQpEuxbxwQXxVLZvTRrUiomqqwbXbGSva8SSHPWh9SU Sintel/Sintel.de.srt
added QmXExs5oDeQwFDN6qRbrrx2JwP2Mho2P4SdoZZNsJyPrF1 Sintel/Sintel.en.srt
added QmPsiMQSVruocq3TG1SmLz2phsBa4XxMadU5f1xyMofHNp Sintel/Sintel.es.srt
added QmdeXx3KWN7LCEECCrT6uHeH6tuvH6eVJRhj5kaDU8LHLS Sintel/Sintel.fr.srt
added QmVovwU6cZXRNCsse4gGibrienX6ZcxWVC9W7FTYeJ2KHh Sintel/Sintel.it.srt
111.8M (117227383B) left
101.5M (106430327B) left
85.1M (89227127B) left
53.7M (56311808B) left
36.2M (37994496B) left
6.2M (6520832B) left
added QmcPZGXSyaQSnDVEtJrgtp2EyxhENiBdpGa8pfs5rMBCWr Sintel/Sintel.mp4
added QmfDZuTjynaCahG7yZrMHBfTccpCHXgs7MHRhFgieLNK5N Sintel/Sintel.nl.srt
added QmPwfmpPuqftfNVnfL5UYVCE5eegSk4jpzxgEtGHnR4oMv Sintel/Sintel.pl.srt
added QmWbquQKiLTPqibijrEeAv8RhRBqDPdehWwNHE3GMVStuH Sintel/Sintel.pt.srt
added QmVDx5LAHNrdej3HXYSqjEX1SYyEtVQ1khqiJ7D8nCG3vP Sintel/Sintel.ru.srt
added QmX5YC15NYxJBzWCWTY6c9PusSefJCTNNZg7VfpndT32aH Sintel/poster.jpg
added QmYSfiNmznRYnSeUQiLVxqzuRwvrJXJLp5ESNb4u5WL1U2
```

Finally, `QmYSfiNmznRYnSeUQiLVxqzuRwvrJXJLp5ESNb4u5WL1U2` is a directory containing
all files from the magnet link.
`

### License 

MIT 2018 - Victor Bjelkholm
