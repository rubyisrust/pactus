package log

import (
	"github.com/pactus-project/pactus/consensus/voteset"
	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/hash"
	"github.com/pactus-project/pactus/types/proposal"
	"github.com/pactus-project/pactus/types/validator"
	"github.com/pactus-project/pactus/types/vote"
)

type Log struct {
	validators    map[crypto.Address]*validator.Validator
	totalPower    int64
	roundMessages []*Messages
}

func NewLog() *Log {
	return &Log{
		roundMessages: make([]*Messages, 0),
	}
}

func (log *Log) RoundMessages(round int16) *Messages {
	return log.mustGetRoundMessages(round)
}

func (log *Log) HasVote(h hash.Hash) bool {
	for _, m := range log.roundMessages {
		if m.HasVote(h) {
			return true
		}
	}
	return false
}

func (log *Log) mustGetRoundMessages(round int16) *Messages {
	for i := int16(len(log.roundMessages)); i <= round; i++ {
		rv := &Messages{
			prepareVotes:   voteset.NewPrepareVoteSet(i, log.totalPower, log.validators),
			precommitVotes: voteset.NewPrecommitVoteSet(i, log.totalPower, log.validators),
			cpPreVotes:     voteset.NewCPPreVoteVoteSet(i, log.totalPower, log.validators),
			cpMainVotes:    voteset.NewCPMainVoteVoteSet(i, log.totalPower, log.validators),
		}

		log.roundMessages = append(log.roundMessages, rv)
	}

	return log.roundMessages[round]
}

func (log *Log) AddVote(v *vote.Vote) (bool, error) {
	m := log.mustGetRoundMessages(v.Round())
	return m.addVote(v)
}

func (log *Log) PrepareVoteSet(round int16) *voteset.BlockVoteSet {
	m := log.mustGetRoundMessages(round)
	return m.prepareVotes
}

func (log *Log) PrecommitVoteSet(round int16) *voteset.BlockVoteSet {
	m := log.mustGetRoundMessages(round)
	return m.precommitVotes
}

func (log *Log) CPPreVoteVoteSet(round int16) *voteset.BinaryVoteSet {
	m := log.mustGetRoundMessages(round)
	return m.cpPreVotes
}

func (log *Log) CPMainVoteVoteSet(round int16) *voteset.BinaryVoteSet {
	m := log.mustGetRoundMessages(round)
	return m.cpMainVotes
}

func (log *Log) HasRoundProposal(round int16) bool {
	return log.RoundProposal(round) != nil
}

func (log *Log) RoundProposal(round int16) *proposal.Proposal {
	m := log.RoundMessages(round)
	if m == nil {
		return nil
	}
	return m.proposal
}

func (log *Log) SetRoundProposal(round int16, proposal *proposal.Proposal) {
	m := log.mustGetRoundMessages(round)
	m.proposal = proposal
}

func (log *Log) MoveToNewHeight(validators []*validator.Validator) {
	log.roundMessages = make([]*Messages, 0)
	log.validators = make(map[crypto.Address]*validator.Validator)
	log.totalPower = 0
	for _, val := range validators {
		log.totalPower += val.Power()
		log.validators[val.Address()] = val
	}
}

func (log *Log) CanVote(addr crypto.Address) bool {
	for _, val := range log.validators {
		if val.Address().EqualsTo(addr) {
			return true
		}
	}
	return false
}
