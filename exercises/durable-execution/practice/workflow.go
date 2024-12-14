package translation

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func SayHelloGoodbye(ctx workflow.Context, input TranslationWorkflowInput) (TranslationWorkflowOutput, error) {
	logger := workflow.GetLogger(ctx)

	logger.Info("SayHelloGoodbye Workflow started", "name", input.Name)

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 45,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	logger.Debug("Executing activity TranslateTerm", "languageCode", input.LanguageCode, "term", "Hello")
	helloInput := TranslationActivityInput{
		Term:         "Hello",
		LanguageCode: input.LanguageCode,
	}
	var helloResult TranslationActivityOutput
	err := workflow.ExecuteActivity(ctx, TranslateTerm, helloInput).Get(ctx, &helloResult)
	if err != nil {
		logger.Error("Failed to translate Hello", "error", err)
		return TranslationWorkflowOutput{}, err
	}
	helloMessage := fmt.Sprintf("%s, %s", helloResult.Translation, input.Name)

	logger.Debug("Sleeping 10 seconds")
	workflow.Sleep(ctx, 10*time.Second)

	logger.Debug("Executing activity TranslateTerm", "languageCode", input.LanguageCode, "term", "Goodbye")
	goodbyeInput := TranslationActivityInput{
		Term:         "Goodbye",
		LanguageCode: input.LanguageCode,
	}
	var goodbyeResult TranslationActivityOutput
	err = workflow.ExecuteActivity(ctx, TranslateTerm, goodbyeInput).Get(ctx, &goodbyeResult)
	if err != nil {
		logger.Error("Failed to translate Goodbye", "error", err)
		return TranslationWorkflowOutput{}, err
	}
	goodbyeMessage := fmt.Sprintf("%s, %s", goodbyeResult.Translation, input.Name)

	output := TranslationWorkflowOutput{
		HelloMessage:   helloMessage,
		GoodbyeMessage: goodbyeMessage,
	}

	return output, nil
}
