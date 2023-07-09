import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { ImageService } from 'src/services/image.service';
import { WasmService } from 'src/services/wasm.service';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss'],
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
    constructor(private readonly _wasmService: WasmService, private readonly _imgService: ImageService, fb: FormBuilder) {
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
