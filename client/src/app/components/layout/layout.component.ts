import { Component, HostListener } from '@angular/core';

@Component({
  selector: 'app-layout',
  templateUrl: './layout.component.html',
  styleUrls: ['./layout.component.scss'],
})
export class LayoutComponent {
  isMobileScreen: boolean = false;
  isDrawerOpen: boolean = true;

  @HostListener('window:resize', ['$event'])
  onWindowResize(event: any) {
    this.isMobileScreen = event.target.innerWidth < 1280;
  }

  toggleDrawer() {
    this.isDrawerOpen = !this.isDrawerOpen;
  }
}
