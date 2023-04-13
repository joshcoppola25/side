# Side

## SIDE - SPOTIFY ID EXTRACTOR

A CLI TOOL FOR EXTRACTING ID FROM SPOTIFY LINKS

**ONLY TESTED ON WINDOWS AND UBUNTU, CLIPBOARD DOES NOT WORK ON UBUNTU**

## WHY?

I have been doing a lot of work with the Spotify API recently and a big chunk of that work has been testing API calls with the Spotify developer console tool they have online.

The issue with that is you need to manually put in song/artist/playlist/album ids.  Spotify used to have a feature where you could copy the ID directly but it seems like they removed that.

Run the program with a link as an argument

Ex. 

```powershell
side [https://open.spotify.com/playlist/2UGCEfh0BKsmWqBx74eN2V?si=964fa37853fb409e](https://open.spotify.com/playlist/2UGCEfh0BKsmWqBx74eN2V?si=964fa37853fb409e)
$: ID is 2UGCEfh0BKsmWqBx74eN2V (It's copied to your clipboard)
```

## INSTALL

Clone the repo to a directory of your choosing

Run:

```powershell
go build
```

Add the folder path to your PATH

Should run without issue
