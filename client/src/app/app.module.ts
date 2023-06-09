// Imports

// Angular
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

// Components
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SigninComponent } from './signin/signin.component';
import { HeaderComponent } from './header/header.component';

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
  declarations: [AppComponent, SigninComponent, HeaderComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    SocialLoginModule,
    GoogleSigninButtonModule,
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
