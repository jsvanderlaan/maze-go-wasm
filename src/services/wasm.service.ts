import { Injectable } from '@angular/core';
import { setupWasm } from '../wasm/wasm_exec.js';

@Injectable({ providedIn: 'root' })
export class WasmService {
    async loadAndRunGoWasm(): Promise<void> {
        setupWasm();
        const go = new Go();
        const result = await WebAssembly.instantiateStreaming(fetch('assets/main.wasm'), go.importObject);
        go.run(result.instance);
    }

    async processImage(blob: Blob, size: number): Promise<string> {
        return URL.createObjectURL(new Blob([await this._getProcessedImage(blob, size)]));
    }

    private _getProcessedImage(file: Blob, size: number): Promise<Uint8Array> {
        return new Promise((resolve, reject) => {
            const reader = new FileReader();

            reader.onload = event => {
                if (event.target === null) {
                    reject();
                    return;
                }

                const bytes = new Uint8Array(event.target.result as ArrayBuffer);
                const threshold = 0.75;

                resolve((window as any).processImage(bytes, threshold, size));
            };

            reader.readAsArrayBuffer(file);
        });
    }
}
