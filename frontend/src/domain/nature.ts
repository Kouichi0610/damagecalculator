/*
    性格補正
*/
export interface INature {
    Name(): string;
    Attack(v: number): number;
    Defense(v: number): number;
    SpAttack(v: number): number;
    SpDefense(v: number): number;
    Speed(v: number): number;
}

// てれや:補正なし
class Bashful implements INature {
    Name(): string { return 'てれや'; }
    Attack(v: number): number { return v; }
    Defense(v: number): number { return v; }
    SpAttack(v: number): number { return v; }
    SpDefense(v: number): number { return v; }
    Speed(v: number): number { return v; }
}
class Lonely extends Bashful {
    Name(): string { return 'さみしがり'; }
    Attack(v: number): number { return Math.floor(v * 1.1); }
    Defense(v: number): number { return Math.floor(v * 0.9); }
}
class Adamant extends Bashful {
    Name(): string { return 'いじっぱり'; }
    Attack(v: number): number { return Math.floor(v * 1.1); }
    SpAttack(v: number): number { return Math.floor(v * 0.9); }
}
class Naughty extends Bashful {
    Name(): string { return 'やんちゃ'; }
    Attack(v: number): number { return Math.floor(v * 1.1); }
    SpDefense(v: number): number { return Math.floor(v * 0.9); }
}
class Brave extends Bashful {
    Name(): string { return 'ゆうかん'; }
    Attack(v: number): number { return Math.floor(v * 1.1); }
    Speed(v: number): number { return Math.floor(v * 0.9); }
}

class Bold extends Bashful {
    Name(): string { return 'ずぶとい'; }
    Defense(v: number): number { return Math.floor(v * 1.1); }
    Attack(v: number): number { return Math.floor(v * 0.9); }
}
class Impish extends Bashful {
    Name(): string { return 'わんぱく'; }
    Defense(v: number): number { return Math.floor(v * 1.1); }
    SpAttack(v: number): number { return Math.floor(v * 0.9); }
}
class Lax extends Bashful {
    Name(): string { return 'のうてんき'; }
    Defense(v: number): number { return Math.floor(v * 1.1); }
    SpDefense(v: number): number { return Math.floor(v * 0.9); }
}
class Relaxed extends Bashful {
    Name(): string { return 'のんき'; }
    Defense(v: number): number { return Math.floor(v * 1.1); }
    Speed(v: number): number { return Math.floor(v * 0.9); }
}

class Modest extends Bashful {
    Name(): string { return 'ひかえめ'; }
    SpAttack(v: number): number { return Math.floor(v * 1.1); }
    Attack(v: number): number { return Math.floor(v * 0.9); }
}
class Mild extends Bashful {
    Name(): string { return 'おっとり'; }
    SpAttack(v: number): number { return Math.floor(v * 1.1); }
    Defense(v: number): number { return Math.floor(v * 0.9); }
}
class Rash extends Bashful {
    Name(): string { return 'うっかりや'; }
    SpAttack(v: number): number { return Math.floor(v * 1.1); }
    SpDefense(v: number): number { return Math.floor(v * 0.9); }
}
class Quiet extends Bashful {
    Name(): string { return 'れいせい'; }
    SpAttack(v: number): number { return Math.floor(v * 1.1); }
    Speed(v: number): number { return Math.floor(v * 0.9); }
}

class Calm extends Bashful {
    Name(): string { return 'おだやか'; }
    SpDefense(v: number): number { return Math.floor(v * 1.1); }
    Attack(v: number): number { return Math.floor(v * 0.9); }
}
class Gentle extends Bashful {
    Name(): string { return 'おとなしい'; }
    SpDefense(v: number): number { return Math.floor(v * 1.1); }
    Defense(v: number): number { return Math.floor(v * 0.9); }
}
class Careful extends Bashful {
    Name(): string { return 'しんちょう'; }
    SpDefense(v: number): number { return Math.floor(v * 1.1); }
    SpAttack(v: number): number { return Math.floor(v * 0.9); }
}
class Sassy extends Bashful {
    Name(): string { return 'なまいき'; }
    SpDefense(v: number): number { return Math.floor(v * 1.1); }
    Speed(v: number): number { return Math.floor(v * 0.9); }
}

class Timid extends Bashful {
    Name(): string { return 'おくびょう'; }
    Speed(v: number): number { return Math.floor(v * 1.1); }
    Attack(v: number): number { return Math.floor(v * 0.9); }
}
class Hasty extends Bashful {
    Name(): string { return 'せっかち'; }
    Speed(v: number): number { return Math.floor(v * 1.1); }
    Defense(v: number): number { return Math.floor(v * 0.9); }
}
class Jolly extends Bashful {
    Name(): string { return 'ようき'; }
    Speed(v: number): number { return Math.floor(v * 1.1); }
    SpAttack(v: number): number { return Math.floor(v * 0.9); }
}
class Naive extends Bashful {
    Name(): string { return 'むじゃき'; }
    Speed(v: number): number { return Math.floor(v * 1.1); }
    SpDefense(v: number): number { return Math.floor(v * 0.9); }
}

let naturesList: INature[] = [
    new Bashful,
    new Lonely,
    new Adamant,
    new Naughty,
    new Brave,
    new Bold,
    new Impish,
    new Lax,
    new Relaxed,
    new Modest,
    new Mild,
    new Rash,
    new Quiet,
    new Calm,
    new Gentle,
    new Careful,
    new Sassy,
    new Timid,
    new Hasty,
    new Jolly,
    new Naive,
];
export function NatureNames(): string[] {
    let res = naturesList.map(n => n.Name());
    return res;
}

export function NatureFactory(name: string): INature {
    // a! undefinedを許容しない
    let res = naturesList.find(n => n.Name() == name)!;
    return res;
}
