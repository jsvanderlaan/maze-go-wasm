import { CommonModule } from '@angular/common';
import { Component, inject, signal } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { filter, map, merge } from 'rxjs';
import { ByteArrayHelper } from 'src/helpers/byte-array.helper';
import { ProcessService } from 'src/services/process.service';
import { ExampleUploadComponent } from './example-upload.component';

@Component({
    imports: [ReactiveFormsModule, CommonModule, ExampleUploadComponent],
    selector: 'app-image-source',
    templateUrl: './image-source.component.html',
})
export class ImageSourceComponent {
    readonly submitted = signal<boolean>(false);

    readonly defaultThreshold = 0.75;
    readonly minThreshold = 0;
    readonly maxThreshold = 1;

    readonly defaultHeight = 80;
    readonly minHeight = 1;
    readonly maxHeight = 1000;

    private processService = inject(ProcessService);
    readonly preview = signal<Uint8Array | null>(null);
    readonly upload$ = merge(toObservable(this.preview), this.processService.sourceImage.pipe(map(input => input.image))).pipe(
        filter(preview => preview !== null),
        map(preview => ByteArrayHelper.toUrl(preview))
    );
    readonly maxSizeMB = 1000;
    readonly exceptedTypes = ['image/jpeg', 'image/png'];

    error: string | null = null;

    readonly form = new FormGroup({
        height: new FormControl(this.defaultHeight, {
            nonNullable: true,
            validators: [Validators.min(this.minHeight), Validators.max(this.maxHeight), Validators.required],
        }),
        threshold: new FormControl(this.defaultThreshold, {
            nonNullable: true,
            validators: [Validators.min(this.minThreshold), Validators.max(this.maxThreshold)],
        }),
        image: new FormControl<Uint8Array | null>(null, {
            validators: [Validators.required],
        }),
    });

    async onSubmit(): Promise<void> {
        this.submitted.set(true);
        if (this.form.valid) {
            this.processService.sourceImage.next(this.form.value as any);
        }
    }

    async onFileSelected(event: Event): Promise<void> {
        const { files } = event.target as HTMLInputElement;

        if (files === null || files.length <= 0) {
            this.form.controls.image.setErrors({ image: 'No file selected' });
            return;
        }

        const file = files[0];
        if (!this.exceptedTypes.includes(file.type)) {
            this.form.controls.image.setErrors({ image: 'Only JPG and PNG images are allowed.' });
            return;
        }
        if (file.size > this.maxSizeMB * 1024 * 1024) {
            this.form.controls.image.setErrors({ image: `File size exceeds ${this.maxSizeMB}MB.` });
            return;
        }

        const arrayBuffer = await file.arrayBuffer();
        const byteArray = new Uint8Array(arrayBuffer);
        this._updateImage(byteArray);
    }

    onPreview(byteArray: Uint8Array): void {
        this._updateImage(byteArray);
    }

    private _updateImage(byteArray: Uint8Array): void {
        this.form.patchValue({
            image: byteArray,
        });
        this.form.controls.image.updateValueAndValidity();
        this.preview.set(byteArray);
    }
}
