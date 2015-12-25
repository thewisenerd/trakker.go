trakker.go
==========

in-memory tracker-list server for p2p

#### what/why?

trouble finding trackers for a torrent?

scrounging the internet for tracker lists?

manually cleaning up your trackers for unresolved host-names, and refused connections?

trakker.go aims to provide a free-to-use server, which helps you overcome all
those problems. how? simple. By providing a service, where _anon_ can add
a tracker to the server, and _anon_ can get a list of all trackers available on
the server.

whenever _anon_ tells trakker.go to add a tracker, it checks if the tracker is
already present, and helps prevent duplicate entries.

whenever _anon_ wants a list of trackers, a utorrent-compatible tracker list
is generated for easy copy-pasting.

#### endpoints

```/list```: lists all trackers

```/add?url=<url>```: add a tracker url

#### todo

* **routing**: add ```/list.txt```, ```/list/genre.txt``` for easy download
* **main**: create a webui for _presentation_, adding trackers
* **trakker**: use JSON for GET/PUT
* **trakker**: add a ```genre``` array field to tracker
* **list**: check if URL is a valid tracker(?) before adding to list
* **list**: ability to sort trackers acc. to ```protcol```, and ```port```
* **list**: use an actual in-memory db (_or build one?_)

---

###### why is this _in-memory_ ???

because it is easier to wipe RAM clean *wink* *wink*
