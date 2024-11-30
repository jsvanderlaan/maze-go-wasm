import { Component } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';

@Component({
    imports: [ReactiveFormsModule],
    selector: 'app-file-upload',
    templateUrl: './file-upload.component.html'
})
export class FileUploadComponent {
    selectedFile: File | null = null;

    onFileSelected(event: Event): void {
        const fileInput = event.target as HTMLInputElement;
        if (fileInput.files && fileInput.files.length > 0) {
            const file = fileInput.files[0];
            if (file.type !== 'image/jpeg') {
                alert('Only JPG images are allowed.');
                return;
            }
            if (file.size > 2 * 1024 * 1024) {
                alert('File size exceeds 2MB.');
                return;
            }
            this.selectedFile = file;
        }
    }

    async onSubmit(event: Event): Promise<void> {
        event.preventDefault();
        if (this.selectedFile) {
            const arrayBuffer = await this.selectedFile.arrayBuffer();
            const byteArray = new Uint8Array(arrayBuffer);
            this.sendToWasm(byteArray);
        }
    }

    sendToWasm(byteArray: Uint8Array): void {
        console.log(byteArray.length);
        // Assuming you have a Go function exposed via WebAssembly
        // if (typeof window.goFunction === 'function') {
        //     window.goFunction(byteArray);
        // } else {
        //     console.error('Go WASM function not found.');
        // }
    }
}
