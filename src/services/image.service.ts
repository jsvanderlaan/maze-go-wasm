import { Injectable } from '@angular/core';

@Injectable({ providedIn: 'root' })
export class ImageService {
    async fetchBlob(url: string): Promise<Blob> {
        const response = await fetch(url);
        return response.blob();
    }
}
