import {HttpClient} from 'aurelia-fetch-client';
import {BACKEND_URL} from '../constants/constants';

export class SelectForm {
  constructor() {
    this.quoteOptions = [
      {'id': 'trump', 'name':'Donald Trump Quotes API'},
      {'id':'taylor', 'name':'Taylor Swift Quotes API'},
      {'id':'kanye', 'name':'Kanye West Quotes API'}
    ];
    this.selectQuoteSource = 'programming';
    this.fetchedQuote = ''
    this.cardDisplay = 'none';
  }


  getFunQuotesData() {
    this.cardDisplay = 'none';
    let httpClient = new HttpClient();
    let urlBase = BACKEND_URL;
    let finalUrl = urlBase + this.selectQuoteSource;
    httpClient.fetch(finalUrl)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        this.fetchedQuote = data.quote;
        this.cardDisplay = 'block';
      });
  }
}
