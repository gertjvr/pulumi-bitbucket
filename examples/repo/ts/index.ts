import * as bitbucket from "@pulumi/bitbucket";

const repo = new bitbucket.Repository("demo-repo", {
  description: "Generated from automated test",
  private: true,
});

export const repoName = repo.name;
