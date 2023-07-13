import { Injectable } from '@angular/core';
import { Theme } from '../enums/Theme';
import { LocalStorageKeys } from 'src/constants/LocalStorageKeys';
import { BehaviorSubject } from 'rxjs';
import { User } from '../interfaces/user.interface';

@Injectable({
  providedIn: 'root',
})
export class LocalStorageService {
  private darkThemeSubject = new BehaviorSubject<boolean>(false);
  public _isDarkTheme = this.darkThemeSubject.asObservable();

  constructor() {
    this.darkThemeSubject.next(this.isDarkTheme());
  }

  // THEME
  setTheme(theme: Theme) {
    localStorage.setItem(LocalStorageKeys.THEME, theme);
    this.darkThemeSubject.next(theme === Theme.DARK);
  }
  setDarkTheme(isDark: boolean) {
    const theme = isDark ? Theme.DARK : Theme.LIGHT;
    localStorage.setItem(LocalStorageKeys.THEME, theme);
    this.darkThemeSubject.next(isDark);
  }
  getTheme(): Theme {
    const themeVal = localStorage.getItem(LocalStorageKeys.THEME);
    const theme = themeVal === Theme.DARK ? Theme.DARK : Theme.LIGHT;
    return theme;
  }
  isDarkTheme(): boolean {
    return this.getTheme() === Theme.DARK;
  }

  // JWT TOKEN
  setAuthToken(token: string) {
    localStorage.setItem(LocalStorageKeys.AUTH_TOKEN, token);
  }
  getAuthToken(): string | null {
    return localStorage.getItem(LocalStorageKeys.AUTH_TOKEN) || null;
  }

  // USER
  setCurrentUser(user: User) {
    localStorage.setItem(LocalStorageKeys.USER, JSON.stringify(user));
  }
  getCurrentUser(): User | null {
    const userVal = localStorage.getItem(LocalStorageKeys.USER);
    return userVal ? (JSON.parse(userVal) as User) : null;
  }
}
