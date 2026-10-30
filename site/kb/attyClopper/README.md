# Ask Attorney Clopper — Knowledge Base

This directory contains plain-text transcripts of Attorney Eric Clopper's
YouTube Shorts, one file per video. They power the "Ask Attorney Clopper"
feature via the standard KB search pipeline.

## File naming

Each file is named by its YouTube video ID:

    {videoID}.txt          e.g.  dQw4w9WgXcQ.txt

## File format

```
# Source: https://www.youtube.com/shorts/{videoID}
# Title: Video title here

Transcript text as a single paragraph of plain text.
```

## Adding transcripts

Run the following command — it reads video IDs from the previously-harvested
YouTube Shorts metadata and fetches the next transcript not yet present here:

    bwctl yt-shorts get

The TranscribeYouTube API (transcribeyoutube.com) is rate limited to
100 requests per day per IP. Run once, commit the result, repeat.

Use `bwctl yt-shorts status` for a progress summary and
`bwctl yt-shorts list` for the full per-video breakdown.

The command is delta-safe: it never overwrites an existing file, so manual
corrections are preserved across runs.
