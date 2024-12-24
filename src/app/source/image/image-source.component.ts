import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { map } from 'rxjs';
import { ByteArrayHelper } from 'src/helpers/byte-array.helper';
import { ProcessService } from 'src/services/process.service';
import { ExampleUploadComponent } from './example-upload.component';

@Component({
    imports: [ReactiveFormsModule, CommonModule, ExampleUploadComponent],
    selector: 'app-image-source',
    templateUrl: './image-source.component.html',
})
export class ImageSourceComponent {
    private processService = inject(ProcessService);
    readonly upload$ = this.processService.sourceImage.pipe(map(arr => ByteArrayHelper.toUrl(arr)));
    readonly maxSizeMB = 1000;
    readonly exceptedTypes = ['image/jpeg', 'image/png'];

    error: string | null = null;

    async onFileSelected(event: Event): Promise<void> {
        const { files } = event.target as HTMLInputElement;

        if (files === null || files.length <= 0) {
            this.error = 'No file selected';
            return;
        }

        const file = files[0];
        if (!this.exceptedTypes.includes(file.type)) {
            this.error = 'Only JPG images are allowed.';
            return;
        }
        if (file.size > this.maxSizeMB * 1024 * 1024) {
            this.error = `File size exceeds ${this.maxSizeMB}MB.`;
            return;
        }

        const arrayBuffer = await file.arrayBuffer();
        const byteArray = new Uint8Array(arrayBuffer);
        this.processService.sourceImage.next(byteArray);
    }
}
