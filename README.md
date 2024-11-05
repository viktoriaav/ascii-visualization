# Ascii Visualization

_Written in_ **GO**
___

Program that takes user inputted data and displays it as ASCII art on a webpage using GET and POST requests.

The base algorithm used for the printing has been reused from the previous ASCII-Art projects.

The user needs to run the server first on their chosen IDE using the ```go run .``` command. 

After that the user needs to navigate to [localhost:8080](https://localhost:8080)

Ascii-art-stylize consists on making your site :
  * more appealing, interactive and intuitive.
  * more user friendly.
  * give more feedback to the user.

___

## How to use

1. Run the server
  * Running: ```go run .```
  * Terminating: Press ```CTRL+C``` in the terminal.
2. Navigate to the [website.](https://localhost:8080)
3. Think of a text you want to print out and the template you want to use.
4. Press **Submit**.

## Details

The program consists of an HTML template that has a dropdown menu where the user can choose the template, a textarea where the user can type in their text and a _submit_ button, to run the program.

The text that user submits is split into words based on the number of line breaks. The final art is created by appending every line of every word, line by line. For example, in a 3 word scenario, the first lines of every word are appended to the final string, then the second lines and so on, until all 8 lines have been appended to the final string.

The final result is then saved to _output.txt_ file.

The art is then displayed on the webpage; it is read line-by-line from the _output.txt_ file.

The _HTML_ files are located under the `/view` folder. All the different ASCII banners are located under the `/fonts` folder. This project does not use any CSS files, every part of visualisation is done in the relevant _HTML_ file.

The GO file: `main.go` starts the server and handles all the GET and POST requests.
When a POST request comes in, the `Ascii` function is started. The inputted string is made into
ASCII-art, sent to the `/ascii-art` handler and displayed.


# Testing:

The program handles errors appropriately and I shall describe how to test them yourself.

* **400 Bad Request** - The server received an invalid input.
* **404 Not Found** - The server cannot find the website which was typed in the address bar.
* **500 Internal Server Error** - The server reached an error state which has not been handled.


## Authors

- [Viktoriia](https://01.kood.tech/git/vavstanc)
- [Anastasiia](https://01.kood.tech/git/akopach)
- [Olha](https://01.kood.tech/git/Olha_Priadkina)