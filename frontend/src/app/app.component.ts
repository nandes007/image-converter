import axios from 'axios'
import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})

export class AppComponent {
  public title:string = 'frontend';

  public async testAxios() {
    try {
      const response = await axios.get('http://localhost:9090');
      console.log(response);
    } catch (error:any) {
      console.log(error);
    }
  }

  async ngOnInit() {
    await this.testAxios()
  }
}
