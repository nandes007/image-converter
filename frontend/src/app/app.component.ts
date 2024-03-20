import axios, { AxiosRequestConfig } from 'axios'
import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import {CommonModule} from '@angular/common';
import { environment } from '../environments/environment';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, CommonModule, FormsModule],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})

export class AppComponent {
  public title:string = "Image Converter";
  public convertTo:string = "";
  public imageUrl:string = "";
  public imageToConvert: File = new File([], 'empty.txt', { type: 'text/plain' });

  public onFileSelected(event:any): void {
    const file: File = event.target.files[0];
    this.imageToConvert = file;
    const reader = new FileReader();

    reader.onload = (e) => {
      this.imageUrl = e.target?.result as string;
    };

    reader.readAsDataURL(file);
  }

  public async doConvert() {
    try {
      const formData = new FormData();
      formData.append('file', this.imageToConvert);
      formData.append('convert_to', this.convertTo);

      const options: AxiosRequestConfig = {
        url: `${environment.apiUrl}/api/v1/process`,
        data: formData,
        method: 'POST'
      }

      const res = await axios(options);
    } catch (error:any) {}
  }

  ngOnInit() {
    console.log(environment);
  }
}
