import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root',
})
export class SigninService {
  constructor(private http: HttpClient) {}

  signin(idToken: string, provider: string) {
    return this.http.post(
      `${environment.API_URL}/signin`,
      {},
      {
        headers: {
          'X-IdToken': idToken,
          'X-Provider': provider,
        },
        observe: 'response',
      }
    );
  }
}
