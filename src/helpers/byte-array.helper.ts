export class ByteArrayHelper {
    static toUrl(arr: Uint8Array): string {
        return URL.createObjectURL(ByteArrayHelper.toBlob(arr));
    }
    static toBlob(arr: Uint8Array, type?: string): Blob {
        return new Blob([arr], { type });
    }
}
