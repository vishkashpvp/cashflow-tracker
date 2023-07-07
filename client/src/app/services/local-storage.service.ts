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
  setAuthToken(token: string) {
    localStorage.setItem(LocalStorageKeys.AUTH_TOKEN, token);
  }
  getAuthToken(): string | null {
    return localStorage.getItem(LocalStorageKeys.AUTH_TOKEN) || null;
  }
}
