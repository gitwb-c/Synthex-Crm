import ForbiddenError from "./models/errors/Forbidden.error";
import UnauthorizedError from "./models/errors/Unauthorized.error";

export default class Helper {
  public static fetchUrl = async (
    url: string,
    init?: RequestInit
  ): Promise<{ data: Record<string, unknown>; statusCode: number }> => {
    const res = await fetch(url, init);
    if (res.status == 401)
      throw new UnauthorizedError("Sessão do usuário expirada!");

    if (res.status == 403)
      throw new ForbiddenError(
        "Usuário não tem permissão para acessar/manipular o conteúdo!"
      );
    const data = await res.json();

    return {
      statusCode: 200,
      data,
    };
  };
}
