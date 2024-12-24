import { NgClass } from '@angular/common';
import { Component, signal } from '@angular/core';
import { SourceType } from 'src/types/source.type';
import { ImageSourceComponent } from './image/image-source.component';
import { TextSourceComponent } from './text/text-source.component';

@Component({
    standalone: true,
    imports: [ImageSourceComponent, TextSourceComponent, NgClass],
    templateUrl: './source.component.html',
    selector: 'app-source',
})
export class SourceComponent {
    readonly selected = signal(SourceType.text);

    source = SourceType;
}
