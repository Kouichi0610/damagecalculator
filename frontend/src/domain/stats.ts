export function calcStats(level, species, individuals, basePoints, nature) {
    let res = [0, 0, 0, 0, 0, 0];
    res[0] = calcHP(level, species[0], individuals[0], basePoints[0]);
    res[1] = nature.attack(calcStat(level, species[1], individuals[1], basePoints[1]));
    res[2] = nature.defense(calcStat(level, species[2], individuals[2], basePoints[2]));
    res[3] = nature.spAttack(calcStat(level, species[3], individuals[3], basePoints[3]));
    res[4] = nature.spDefense(calcStat(level, species[4], individuals[4], basePoints[4]));
    res[5] = nature.speed(calcStat(level, species[5], individuals[5], basePoints[5]));
    return res;
}

function calcHP(l, s, i, b) {
    let x = s * 2 + i + b/4;
    let y = x * l / 100.0 + l + 10;
    return Math.floor(y);
}
function calcStat(l, s, i, b) {
  let x = s * 2 + i + b/4;
  let y = Math.floor(x * l / 100.0 + 5);
  return Math.floor(y);
}
