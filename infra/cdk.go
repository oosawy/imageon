package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	awscdklambdago "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkStackProps struct {
	awscdk.StackProps
}

func NewCdkStack(scope constructs.Construct, id string, props *CdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	unprocessedBucket := awss3.NewBucket(stack, jsii.String("unprocessed-images-bucket"), &awss3.BucketProps{
		LifecycleRules: &[]*awss3.LifecycleRule{
			{
				Expiration: awscdk.Duration(awscdk.Duration_Days(jsii.Number(30))),
			},
		},
	})

	processedBucket := awss3.NewBucket(stack, jsii.String("processed-images-bucket"), &awss3.BucketProps{})
	publicBucketPolicy := awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect: awsiam.Effect_ALLOW,
		Actions: &[]*string{
			jsii.String("s3:GetObject"),
		},
		Principals: &[]awsiam.IPrincipal{
			awsiam.NewAnyPrincipal(),
		},
		Resources: &[]*string{
			jsii.String(*unprocessedBucket.BucketArn() + "/*"),
		},
	})
	processedBucket.AddToResourcePolicy(publicBucketPolicy)

	lambda := awscdklambdago.NewGoFunction(stack, jsii.String("handler"), &awscdklambdago.GoFunctionProps{
		Entry: jsii.String("../lambda"),
		Environment: &map[string]*string{
			"UNPROCESSED_BUCKET": unprocessedBucket.BucketArn(),
			"PROCESSED_BUCKET":   processedBucket.BucketArn(),
		},
	})

	allowedOrigin := "*"
	lambda.AddFunctionUrl(&awslambda.FunctionUrlOptions{
		AuthType: awslambda.FunctionUrlAuthType_NONE,
		Cors: &awslambda.FunctionUrlCorsOptions{
			AllowedMethods: &[]awslambda.HttpMethod{awslambda.HttpMethod_ALL},
			AllowedOrigins: &[]*string{&allowedOrigin},
		},
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkStack(app, "CdkStack", &CdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
