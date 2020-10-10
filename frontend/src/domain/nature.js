// 性格
class nature {
    constructor(at, df, sa, sd, sp, name) {
        this.at = at;
        this.df = df;
        this.sa = sa;
        this.sd = sd;
        this.sp = sp;
        this.name = name;
    }
    attack(val) {
        return Math.floor(val * this.at);
    }
    defense(val) {
        return Math.floor(val * this.df);
    }
    spAttack(val) {
        return Math.floor(val * this.sa);
    }
    spDefense(val) {
        return Math.floor(val * this.sd);
    }
    speed(val) {
        return Math.floor(val * this.sp);
    }
}
// 性格の名前一覧
export function natureNames() {
    return natures.map(n => n.name);
}
// 名前から性格を取得
export function getNature(name) {
    return natures.find(n => n.name == name);
}

let natures = [];
natures.push(new nature(1.0, 1.0, 1.0, 1.0, 1.0, 'てれや'))

natures.push(new nature(1.1, 0.9, 1.0, 1.0, 1.0, 'さみしがり'))
natures.push(new nature(1.1, 1.0, 0.9, 1.0, 1.0, 'いじっぱり'))
natures.push(new nature(1.1, 1.0, 1.0, 0.9, 1.0, 'やんちゃ'))
natures.push(new nature(1.1, 1.0, 1.0, 1.0, 0.9, 'ゆうかん'))

natures.push(new nature(0.9, 1.1, 1.0, 1.0, 1.0, 'ずぶとい'))
natures.push(new nature(1.0, 1.1, 0.9, 1.0, 1.0, 'わんぱく'))
natures.push(new nature(1.0, 1.1, 1.0, 0.9, 1.0, 'のうてんき'))
natures.push(new nature(1.0, 1.1, 1.0, 1.0, 0.9, 'のんき'))

natures.push(new nature(0.9, 1.0, 1.1, 1.0, 1.0, 'ひかえめ'))
natures.push(new nature(1.0, 0.9, 1.1, 1.0, 1.0, 'おっとり'))
natures.push(new nature(1.0, 1.0, 1.1, 0.9, 1.0, 'うっかりや'))
natures.push(new nature(1.0, 1.0, 1.1, 1.0, 0.9, 'れいせい'))

natures.push(new nature(0.9, 1.0, 1.0, 1.1, 1.0, 'おだやか'))
natures.push(new nature(1.0, 0.9, 1.0, 1.1, 1.0, 'おとなしい'))
natures.push(new nature(1.0, 1.0, 0.9, 1.1, 1.0, 'しんちょう'))
natures.push(new nature(1.0, 1.0, 1.0, 1.1, 0.9, 'なまいき'))

natures.push(new nature(0.9, 1.0, 1.0, 1.0, 1.1, 'おくびょう'))
natures.push(new nature(1.0, 0.9, 1.0, 1.0, 1.1, 'せっかち'))
natures.push(new nature(1.0, 1.0, 0.9, 1.0, 1.1, 'ようき'))
natures.push(new nature(1.0, 1.0, 1.0, 0.9, 1.1, 'むじゃき'))
