import { Component, Renderer2 } from '@angular/core';
import { Theme } from './enum/theme';
import { LocalStorageService } from './services/local-storage.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  title = 'cashflow-tracker-client';
  isDarkTheme: boolean = false;
  html: HTMLElement;

  constructor(
    private renderer: Renderer2,
    private localStorageService: LocalStorageService
  ) {
    this.html = document.documentElement;
    this.isDarkTheme = this.localStorageService.getTheme() === Theme.DARK;
    this.applyTheme();
  }

  toggleTheme() {
    this.isDarkTheme = !this.isDarkTheme;
    this.applyTheme();
    this.localStorageService.setTheme(
      this.isDarkTheme ? Theme.DARK : Theme.LIGHT
    );
  }

  applyTheme() {
    this.isDarkTheme
      ? this.renderer.addClass(this.html, Theme.DARK)
      : this.renderer.removeClass(this.html, Theme.DARK);
  }
}
