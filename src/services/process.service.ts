import { Injectable } from '@angular/core';
import { map, Observable, Subject } from 'rxjs';

@Injectable({
    providedIn: 'root',
})
export class ProcessService {
    private _original: Subject<Uint8Array> = new Subject();
    readonly output: Observable<Uint8Array>;

    constructor() {
        this.output = this._original.pipe(map(arr => arr));
    }

    setOriginal(arr: Uint8Array): void {
        this._original.next(arr);
    }
}
