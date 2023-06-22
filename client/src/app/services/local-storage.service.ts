import { Injectable } from '@angular/core';
import { Theme } from '../enum/theme';
import { LocalStorageKeys } from 'src/constants/LocalStorageKeys';

@Injectable({
  providedIn: 'root',
})
export class LocalStorageService {
  constructor() {}

  setTheme(theme: Theme) {
    localStorage.setItem(LocalStorageKeys.Theme, theme);
  }

  getTheme(): Theme {
    const theme = localStorage.getItem(LocalStorageKeys.Theme);
    return theme ? (theme as Theme) : Theme.LIGHT;
  }
}
