import {Component, OnInit} from '@angular/core';
import {HelloWorldService} from './hello-world.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {

  title;

  constructor(private hw: HelloWorldService) {}

  ngOnInit() {
    this.hw.getTitle()
      .subscribe(data => this.title = data.title);

    console.log(this.title);
  }

}
