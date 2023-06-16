// Imports

// Angular
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatIconModule } from '@angular/material/icon';

// Components
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SigninComponent } from './signin/signin.component';
import { HeaderComponent } from './header/header.component';
import { HomeComponent } from './home/home.component';

import { environment } from 'src/environments/environment';

// External
import {
  GoogleLoginProvider,
  GoogleSigninButtonModule,
  SocialAuthServiceConfig,
  SocialLoginModule,
} from '@abacritt/angularx-social-login';

const GOOGLE = {
  id: GoogleLoginProvider.PROVIDER_ID,
  provider: new GoogleLoginProvider(environment.G_CLIENT_ID),
};

const socialAuthServiceConfig: SocialAuthServiceConfig = {
  autoLogin: false,
  providers: [GOOGLE],
  onError: (err) => {
    console.error(err);
  },
};

@NgModule({
  declarations: [AppComponent, SigninComponent, HeaderComponent, HomeComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    SocialLoginModule,
    GoogleSigninButtonModule,
    BrowserAnimationsModule,
    MatIconModule,
  ],
  providers: [
    {
      provide: 'SocialAuthServiceConfig',
      useValue: socialAuthServiceConfig,
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
