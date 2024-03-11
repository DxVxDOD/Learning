export default function two_crystal_balls(breaks: boolean[]): number {
    const jumpAmoumt = Math.floor(Math.sqrt(breaks.length));
    let i = jumpAmoumt;
    for (; i < breaks.length; i += jumpAmoumt) {
        if (breaks[i]) {
            break;
        }
    }
    i -= jumpAmoumt;
    for (let j = 0; j < jumpAmoumt && i < breaks.length; ++j, ++i) {
        if (breaks[i]) {
            return i;
        }
    }

    return -1;
}
