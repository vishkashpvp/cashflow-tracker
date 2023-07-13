import { Component, OnInit, Renderer2 } from '@angular/core';
import { Theme } from './enums/Theme';
import { LocalStorageService } from './services/local-storage.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit {
  title = 'cashflow-tracker-client';
  isDarkTheme: boolean = false;
  html: HTMLElement;

  constructor(
    private renderer: Renderer2,
    private localStorageService: LocalStorageService
  ) {
    this.html = document.documentElement;
    this.isDarkTheme = this.localStorageService.isDarkTheme();
    this.applyTheme();
  }

  ngOnInit(): void {
    this.localStorageService._isDarkTheme.subscribe((_isDarkTheme: boolean) => {
      this.isDarkTheme = _isDarkTheme;
      this.applyTheme();
    });
  }

  toggleTheme() {
    this.isDarkTheme = !this.isDarkTheme;
    this.applyTheme();
    this.localStorageService.setDarkTheme(this.isDarkTheme);
  }

  applyTheme() {
    this.isDarkTheme
      ? this.renderer.addClass(this.html, Theme.DARK)
      : this.renderer.removeClass(this.html, Theme.DARK);
  }
}
