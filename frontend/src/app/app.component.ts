import axios from 'axios'
import { Component, NgModule } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import {CommonModule} from '@angular/common';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, CommonModule],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})

export class AppComponent {
  public title:string = "Image Converter";
  public convertTo:string = "";
  public imageUrl:string = "";

  public onFileSelected(event:any): void {
    const file: File = event.target.files[0];
    const reader = new FileReader();

    reader.onload = (e) => {
      this.imageUrl = e.target?.result as string;
    };

    reader.readAsDataURL(file);
  }

  ngOnInit() {
  }
}
