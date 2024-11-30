import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { map } from 'rxjs';
import { ByteArrayHelper } from 'src/helpers/byte-array.helper';
import { ProcessService } from 'src/services/process.service';

@Component({
    imports: [ReactiveFormsModule, CommonModule],
    selector: 'app-file-upload',
    templateUrl: './file-upload.component.html',
})
export class FileUploadComponent {
    private processService = inject(ProcessService);
    readonly upload$ = this.processService.original.pipe(map(arr => ByteArrayHelper.toUrl(arr)));
    readonly maxSizeMB = 2;
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
        this.processService.original.next(byteArray);
    }
}
