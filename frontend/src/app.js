import {HttpClient} from 'aurelia-fetch-client';
import {PLATFORM} from 'aurelia-pal';


export class App {
  constructor() {
  }

  configureRouter(config, router) {
      config.title = '';

      config.map([
        { route: ['','/'],  name: 'home',
           moduleId: './selectform/selectform',  nav: true, title:'Home' },

         { route: 'fun',  name: 'funform',
            moduleId: PLATFORM.moduleName('./funform/funform'),  nav: true, title:'Fun' }
      ]);

      this.router = router;
   }
}
