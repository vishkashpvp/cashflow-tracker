import { Component } from '@angular/core';
import { Theme } from 'src/app/enums/Theme';
import { LocalStorageService } from 'src/app/services/local-storage.service';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.scss'],
})
export class SettingsComponent {
  selectedTheme: Theme;
  selectedThemeOption: string = 'light';

  constructor(private localStorageService: LocalStorageService) {
    this.selectedTheme = this.localStorageService.getTheme();
    this.selectedThemeOption = this.selectedTheme.toLocaleLowerCase();
  }

  themeToggle() {
    const isDark = this.selectedThemeOption.toLowerCase() === Theme.DARK;
    this.localStorageService.setTheme(isDark ? Theme.DARK : Theme.LIGHT);
  }
}
