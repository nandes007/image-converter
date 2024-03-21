import axios from 'axios'
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import {CommonModule} from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { environment } from '../environments/environment';
import { LoadingComponent } from './components/loading/loading.component';
import { ErrorMessageComponent } from './components/error-message/error-message.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    RouterOutlet,
    CommonModule,
    FormsModule,
    ErrorMessageComponent,
    LoadingComponent
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})

export class AppComponent {
  public convertTo:string = "";
  public imageUrl:string = "";
  public imageToConvert: File = new File([], 'empty.txt', { type: 'text/plain' });
  public errorMessage:string = "";
  public isSubmit:boolean = false;

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
    this.isSubmit = true;
    try {
      const formData = new FormData();
      formData.append('file', this.imageToConvert);
      formData.append('convert_to', this.convertTo);
      const response = await axios.post(`${environment.apiUrl}/api/v1/process`, formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      });

      // Extract filename from Content-Disposition header
      let filename = 'downloaded-file';
      const contentDispositionHeader = response.headers['content-disposition'];
      if (contentDispositionHeader) {
          const filenameRegex = /filename[^;=\n]*=((['"]).*?\2|[^;\n]*)/;
          const matches = filenameRegex.exec(contentDispositionHeader);
          if (matches && matches.length > 1) {
              filename = decodeURIComponent(matches[1].replace(/['"]/g, ''));
          }
      }

      // Create a blob URL and initiate the file download
      const blob = new Blob([response.data], { type: response.headers['content-type'] });
      const url = window.URL.createObjectURL(blob);
      const link = document.createElement('a');
      link.href = url;
      link.setAttribute('download', filename);
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
      window.URL.revokeObjectURL(url);
      
      this.isSubmit = false;
    } catch (error:any) {
      this.isSubmit = false;
      this.errorMessage = error.response.data.message;
      setTimeout(() => {
        this.errorMessage = "";
      }, 3000);
    }
  }

  ngOnInit() {
  }
}
