export default interface PipelineContract {
  pipelines: {
    edges: {
      node: {
        id: string;
        name: string;
      }[];
    };
  };
}
