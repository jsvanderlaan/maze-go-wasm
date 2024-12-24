import { Injectable } from '@angular/core';
import { Settings } from 'src/app/settings.component.js';

@Injectable({ providedIn: 'root' })
export class WorkerService {
    private _worker: Worker | null = null;

    processImage(byteArray: Uint8Array, { threshold, size }: Settings): Promise<any> {
        return this._execute('processImage', { byteArray, threshold, size });
    }

    processText(text: string, outline: boolean, { threshold, size }: Settings): Promise<any> {
        return this._execute('processText', { text, outline, threshold, size });
    }

    private _execute(command: string, payload: any): Promise<any> {
        const worker = this._createWorker();

        return new Promise((resolve, reject) => {
            if (worker === null) {
                reject(new Error('worker should not be null at this point'));
                return;
            }

            worker.onmessage = ({ data }) => {
                const { status, result, error } = data;
                if (status === 'completed') {
                    resolve(result);
                } else if (status === 'error') {
                    reject(new Error(error));
                }
            };

            worker.postMessage({ command, payload });
        });
    }

    private _createWorker(): Worker {
        if (this._worker !== null) {
            this._worker.terminate();
        }
        this._worker = new Worker(new URL('../workers/process.worker', import.meta.url), { type: 'module' });
        return this._worker;
    }
}
