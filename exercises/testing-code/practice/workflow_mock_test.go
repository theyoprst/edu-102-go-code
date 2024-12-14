package translation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/testsuite"
)

func TestSuccessfulTranslationWithMocks(t *testing.T) {
	s := testsuite.WorkflowTestSuite{}

	env := s.NewTestWorkflowEnvironment()

	workflowInput := TranslationWorkflowInput{
		Name:         "Pierre",
		LanguageCode: "fr",
	}

	helloInput := TranslationActivityInput{Term: "Hello", LanguageCode: "fr"}
	helloOutput := TranslationActivityOutput{Translation: "Bonjour"}
	env.OnActivity(TranslateTerm, mock.Anything, helloInput).Return(helloOutput, nil)

	goodbyeInput := TranslationActivityInput{Term: "Goodbye", LanguageCode: "fr"}
	goodgyeOutput := TranslationActivityOutput{Translation: "Au revoir"}
	env.OnActivity(TranslateTerm, mock.Anything, goodbyeInput).Return(goodgyeOutput, nil)

	env.ExecuteWorkflow(SayHelloGoodbye, workflowInput)

	assert.True(t, env.IsWorkflowCompleted())

	var result TranslationWorkflowOutput
	assert.NoError(t, env.GetWorkflowResult(&result))
	assert.Equal(t, result.HelloMessage, "Bonjour, Pierre")
	assert.Equal(t, result.GoodbyeMessage, "Au revoir, Pierre")
}
