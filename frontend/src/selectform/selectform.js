import {HttpClient} from 'aurelia-fetch-client';
import {getRandomInt} from '../utilities/numutils'
import {BACKEND_URL} from '../constants/constants';

export class SelectForm {
  constructor() {
    this.quoteOptions = [
      {'id': 'programming', 'name':'Programming Quotes API'},
      {'id':'garden', 'name':'Garden Quotes API'},
      {'id':'design', 'name':'Design Quotes API'},
      {'id':'favorite', 'name': 'Favorite Quotes API'},
      {'id':'forismatic', 'name': 'Forismatic Quotes API'}
    ];
    this.selectQuoteSource = 'programming';
    this.fetchedQuote = ''
    this.cardDisplay = 'none';
    this.sentimentDisplay = 'none';
    this.sentimentOptions = [
      { 'title': 'Your Awesome Sentiment Card!',
        'subtitle': 'Positive',
        'bartype': 'success'},
      { 'title': 'Your Awesome Sentiment Card!',
        'subtitle': 'Negative',
        'bartype': 'warning'
      },
      {'title': 'Your Awesome Sentiment Card!',
        'subtitle': 'Neutral',
        'bartype': 'primary'
      }
    ];
  }


  getQuotesData() {
    this.cardDisplay = 'none';
    this.sentimentDisplay = 'none';
    let httpClient = new HttpClient();
    let urlBase = BACKEND_URL;
    let finalUrl = urlBase + this.selectQuoteSource;
    httpClient.fetch(finalUrl)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        this.fetchedQuote = data.quote;
        this.cardDisplay = 'block';
        for(let i=0; i < this.sentimentOptions.length; i++) {
          this.sentimentOptions[i].weight = getRandomInt(100);
        }
        this.sentimentDisplay = 'block';
      });
  }
}
