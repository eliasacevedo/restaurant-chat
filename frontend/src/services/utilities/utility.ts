export interface IUtility {
    addZeros: (value: number, length: number) => string;
}

export class UtilityServices implements IUtility {
    addZeros(value: number, length: number) {
        return `${'0'.repeat(length - value.toString().length)}${value}`
    }

}