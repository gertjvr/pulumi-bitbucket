import pulumi
import pulumi_bitbucket as bitbucket

repo = bitbucket.Repository("demo-repo",
    description="Generated from automated test",
    private="true"
)

pulumi.export('repo_name', repo.name)


