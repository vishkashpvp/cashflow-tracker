import { Injectable } from '@angular/core';
import { Theme } from '../enum/theme';
import { LocalStorageKeys } from 'src/constants/LocalStorageKeys';

@Injectable({
  providedIn: 'root',
})
export class LocalStorageService {
  constructor() {}

  // THEME
  setTheme(theme: Theme) {
    localStorage.setItem(LocalStorageKeys.THEME, theme);
  }
  getTheme(): Theme {
    const theme = localStorage.getItem(LocalStorageKeys.THEME);
    return theme ? (theme as Theme) : Theme.LIGHT;
  }

  // JWT TOKEN
  setToken(token: string) {
    localStorage.setItem(LocalStorageKeys.TOKEN, token);
  }
  getToken(): string | null {
    return localStorage.getItem(LocalStorageKeys.TOKEN) || null;
  }
}
