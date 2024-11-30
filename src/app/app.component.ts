import { NgIf } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { ImageService } from 'src/services/image.service';
import { ProcessService } from 'src/services/process.service';
import { WasmService } from 'src/services/wasm.service';
import { FileUploadComponent } from './file-upload.component';
import { ResultComponent } from './result.component';

@Component({
    standalone: true,
    imports: [ReactiveFormsModule, FileUploadComponent, ResultComponent, NgIf],
    selector: 'app-root',
    templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {
    readonly sizeMin: number = 10;
    readonly sizeMax: number = 400;
    resultSrc: string = '';
    loading: boolean = false;

    form: FormGroup;
    controls = {
        size: new FormControl(100, [Validators.min(this.sizeMin), Validators.max(this.sizeMax)]),
    };
    constructor(
        private readonly _wasmService: WasmService,
        private readonly _imgService: ImageService,
        processService: ProcessService
    ) {
        this.form = new FormGroup(this.controls);
    }

    async ngOnInit(): Promise<void> {
        await this._wasmService.loadAndRunGoWasm();
    }

    async submit(): Promise<void> {
        this.loading = true;
        const size = this.form.value.size;
        console.log(size);
        const blob = await this._imgService.fetchBlob('assets/test.jpg');
        const url = await this._wasmService.processImage(blob, size);
        this.resultSrc = url;
        this.loading = false;
    }
}
