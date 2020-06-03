<h1 align="center">XKCD Comic Service</h1>
<p>
    <img src='https://img.shields.io/badge/maintainer-harshit98-blue' />
  <img src="https://img.shields.io/badge/license-MIT-blue" />
  <a href="https://github.com/itsfadnis/coronavirus-india/graphs/commit-activity" target="_blank">
    <img alt="Maintenance" src="https://img.shields.io/badge/Maintained%3F-yes-green.svg" />
  </a>
  <a href="https://twitter.com/HarshitPrasad8" target="_blank">
    <img alt="Twitter: HarshitPrasad8" src="https://img.shields.io/twitter/follow/HarshitPrasad8.svg?style=social" />
  </a>
</p>

> Meet Comic Ghost. He is a friendly ghost which fetches web comics for you from xkcd platform using CLI and downloads them as an image (only if you want!).
> Just order him and he will serve you well.

## 🎥 Demo Video

Link - https://youtu.be/paOWKCzcMl4

## Prerequisites

- Go 1.14

## Features

1. If you run the program without any command line args, it will fetch latest comic book.
2. You can fetch comic book according to comic book number.
3. Provided an option if user wants to save the comic. By default, it does not save the comic.
4. Option to print response from API in JSON format.

## Usage

- Open your terminal and `cd` into the project repo.

- Build `main.go`
  ```
  go build -o ghost-bring-me-comicbook  
  ```

- Run the executable binary file
  ```
  ./ghost-bring-me-comicbook --num=324 --save=true --output=json 
  ```

- If you enabled `--save=true` flag, then you can find the downloaded comic within same directory.

## Author

👤 **Harshit Prasad**

* Twitter: [@HarshitPrasad8](https://twitter.com/HarshitPrasad8)
* Github: [@harshit98](https://github.com/harshit98)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!

## Show your support

Give a ⭐️ if you think this project is awesome!

## 📝 License

Copyright © 2020 [Harshit Prasad](https://github.com/harshit98).<br />
This project is [MIT](https://github.com/harshit98/Comic-Ghost/blob/master/LICENSE) licensed.
