import { Component, Renderer2 } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  title = 'cashflow-tracker-client';
  isDarkTheme: boolean = false;
  html: HTMLElement;

  constructor(private render: Renderer2) {
    this.html = document.documentElement;
    this.render.removeClass(this.html, 'dark');
  }

  toggleTheme() {
    this.isDarkTheme = !this.isDarkTheme;

    this.isDarkTheme
      ? this.render.addClass(this.html, 'dark')
      : this.render.removeClass(this.html, 'dark');
  }
}
