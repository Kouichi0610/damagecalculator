import { BasePoints } from '../../src/store/basePoints/types'
// TODO:@/store/

describe('基礎ポイント', (): void => {
  test('割り振り確認', (): void => {
    var basePoints = new BasePoints();

    // 初期状態では252まで
    expect(basePoints.defense().limit()).toBe(252);
    // 0～252であれば任意の値を設定できる
    basePoints.set(0, -1);
    expect(basePoints.hp().value).toBe(0);
    basePoints.set(0, 80);
    expect(basePoints.hp().value).toBe(80);
    basePoints.set(1, 253);
    expect(basePoints.attack().value).toBe(252);

    // 合計508まで
    expect(basePoints.defense().limit()).toBe(176);
    basePoints.set(2, 4);
    expect(basePoints.defense().value).toBe(4);
    basePoints.set(2, 177);
    expect(basePoints.defense().value).toBe(176);
  })
})
