export default class Helper {
  public static getContrastColor = (
    color: string | number[] | { r: number; g: number; b: number; a?: number }
  ): string => {
    let r: number, g: number, b: number, a: number;

    if (typeof color === "string") {
      const rgbaMatch = color.match(
        /rgba?\((\d+),\s*(\d+),\s*(\d+)(?:,\s*([\d.]+))?\)/i
      );
      if (rgbaMatch) {
        r = parseInt(rgbaMatch[1]);
        g = parseInt(rgbaMatch[2]);
        b = parseInt(rgbaMatch[3]);
        a = rgbaMatch[4] ? parseFloat(rgbaMatch[4]) : 1;
      } else {
        const numbers = color.match(/(\d+)/g);
        if (numbers && numbers.length >= 3) {
          r = parseInt(numbers[0]);
          g = parseInt(numbers[1]);
          b = parseInt(numbers[2]);
          a = numbers[3] ? parseFloat(numbers[3]) : 1;
        } else {
          return "black"; // fallback
        }
      }
    } else if (Array.isArray(color)) {
      [r, g, b] = color;
      a = color[3] !== undefined ? color[3] : 1;
    } else {
      r = color.r;
      g = color.g;
      b = color.b;
      a = color.a !== undefined ? color.a : 1;
    }

    if (a < 1) {
      r = Math.round(r * a + 255 * (1 - a));
      g = Math.round(g * a + 255 * (1 - a));
      b = Math.round(b * a + 255 * (1 - a));
    }

    const luminance = (0.299 * r + 0.587 * g + 0.114 * b) / 255;

    return luminance > 0.5 ? "black" : "white";
  };
}
