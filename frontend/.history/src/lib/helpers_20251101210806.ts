export default class Helper {
  public static getContrastColor = (color: string) => {
    let r, g, b, a;

    if (typeof color === "string") {
      // Suporte para rgba(r, g, b, a)
      const rgbaMatch = color.match(
        /rgba?\((\d+),\s*(\d+),\s*(\d+)(?:,\s*([\d.]+))?\)/i
      );
      if (rgbaMatch) {
        r = parseInt(rgbaMatch[1]);
        g = parseInt(rgbaMatch[2]);
        b = parseInt(rgbaMatch[3]);
        a = rgbaMatch[4] ? parseFloat(rgbaMatch[4]) : 1; // default alpha = 1
      } else {
        // Suporte para "r, g, b, a" ou "r, g, b"
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
      // Se for array [r, g, b] ou [r, g, b, a]
      [r, g, b] = color;
      a = color[3] !== undefined ? color[3] : 1;
    } else if (typeof color === "object") {
      // Se for objeto {r, g, b} ou {r, g, b, a}
      r = color.r;
      g = color.g;
      b = color.b;
      a = color.a !== undefined ? color.a : 1;
    }

    // Se tiver alpha, considera o fundo branco para cálculo do contraste
    if (a < 1) {
      // Blend com fundo branco (assumindo fundo branco)
      r = Math.round(r * a + 255 * (1 - a));
      g = Math.round(g * a + 255 * (1 - a));
      b = Math.round(b * a + 255 * (1 - a));
    }

    // Calcula a luminância relativa (fórmula WCAG)
    const luminance = (0.299 * r + 0.587 * g + 0.114 * b) / 255;

    // Retorna preto para fundos claros, branco para fundos escuros
    return luminance > 0.5 ? "black" : "white";
  };
}
