# coub-backup

# Table of Contents
1. [About](#about)
2. [Usage](#usage)
3. [Contributing](#contributing)
4. [License](#license)

## About
coub-backup is a simple tool to automate the downloading of coubs from [coub.com](https://coub.com).

This tool will parse a user profile for all of their coubs, and assemble directory structure by year and month for each coub.

It will generate an info text file that looks like:

```text
Title: Atrium - Week End
Created At: 2021-05-31 04:37:17 +0000 UTC
Duration: 8.44
Views: 2754
Recoubs: 4
Source: map[has_embed:true service_name:YouTube type:Youtube url:https://youtu.be/TUYmmvFPEro]
Tags: synthesizer, synthwave, synth, 80s dancing, dancing, classics, live, 80s music, eighties, 80s, disco, italodisco, italo, 80, palpala, italo disco, guillermo gustavo gallardo, guille music video, week end, atrium

```

Additional metadata is also generated for each coub, including all original coub information. This metadata is stored in a JSON file. 

## Usage:

If building from source, with Golang >v.1.17.0, use `go build` in the main directory to get coub-backup.exe.

### Backup a user's coubs

`coub-backup -directory=<directory> <username>`

### Backup a community's coubs (max of 500 pages)

`coub-backup -directory=<directory> -community=<community> -pages=10`

Supported options: *animals-pets, mashup, music, blogging, standup-jokes, movies, anime, gaming, cartoons, art, live-pictures, news, sports, science-technology, food-kitchen, celebrity, nature-travel, fashion, dance, cars, memes, nsfw*

### Backup a Best Of for a given year (between 2012 and 2022)

`coub-backup -directory=<directory> -bestof=<year>`

### Backup the day's coub of the day

`coub-backup -directory=<directory> -day=true`

### Backup the day's featured coubs

`coub-backup -directory=<directory> -featured=true`

## Contributing

It's highly recommended to join the discord group before submitting issues or pull requests: [https://discord.gg/Z7WYmbaGpU](https://discord.gg/Z7WYmbaGpU)

## License

Unless otherwise specified, all content is licensed under the [GPLV3](https://www.gnu.org/licenses/gpl-3.0.en.html) license.