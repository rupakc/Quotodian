# Quotodian
A consolidated platform for fetching and displaying quotes

## System architecture
![architecture image][logo]

[logo]: https://github.com/rupakc/Quotodian/blob/main/frontend/static/quotodian-architecture.png
<p>The overarching aims of Quotodian is to provide a consolidated platform for browsing through quotes.
                    It fetches random quotes on a wide variety of topics like Programming, Design, Philosophy etc.
                    The backend has been developed in <b> Golang </b> and the front-end uses <b> Aurelia JS </b> framework, along with <b> PaperCSS </b>
                    There is also a tab dedicated for fun which contains tweets of <b> Donald Trump, Taylor Swift </b> and <b> Kanye West </b>
                  </p>

## Steps to run the Application

  - Clone the repository
  - Install Golang
  - Install Node (version 10 and above)
  - Install Aurelia JS
  - cd into the Backend directory and run the backend server by executing the command - ``` go run quotesapi.go ```
  - cd into the Frontend directory and execute the following commands - 
      - ``` npm install ```
      - ``` au run --open ```
