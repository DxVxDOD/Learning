export default function bfs(head: BinaryNode<number>, needle: number): boolean {
    const q: (BinaryNode<number> | null)[] = [head];

    while (q.length) {
        const curr = q.shift() as BinaryNode<number> | null;

        if (!curr) {
            continue;
        }

        if (curr.value === needle) {
            return true;
        }
        q.unshift(curr.left);
        q.unshift(curr.right);
    }
    return false;
}
