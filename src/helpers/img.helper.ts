export class ImgHelper {
    static toUrl(arr: Uint8Array): string {
        return URL.createObjectURL(new Blob([arr]));
    }
}
