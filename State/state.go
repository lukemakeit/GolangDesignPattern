package State

import "fmt"

//封装一个投票状态相关的行为
type VoteSate interface {
	Vote(user string,voteItem string,voteManager *VoteManager)
}
//正常投票
type NormalVoteState struct {}
func (nvs *NormalVoteState)Vote(user string,voteItem string,voteManager *VoteManager)  {
	//记录到投票结果
	voteManager.MapVote[user]=voteItem
	fmt.Printf("%s 恭喜你投票成功\n",user)
}
//对于投票成功的用户,额外再给予10个积分做奖励
type NormalVoteState2 struct {
	*NormalVoteState
}
func (nvs2 *NormalVoteState2)Vote(user string,voteItem string,voteManager *VoteManager)  {
	nvs2.NormalVoteState.Vote(user,voteItem,voteManager)
	fmt.Printf("%s 请查收10个积分\n",user)
}
//重复投票
type RepeatVoteState struct {}
func (rvs *RepeatVoteState)Vote(user string,voteItem string,voteManager *VoteManager)  {
	fmt.Printf("%s 请不要重复投票\n",user)
}
//恶意投票
type SpiteVoteState struct {}
func (svs *SpiteVoteState)Vote(user string,voteItem string,voteManager *VoteManager)  {
	//取消用户投票资格,并取消投票记录
	_,ok:=voteManager.MapVote[user]
	if ok==true {
		delete(voteManager.MapVote,user)
	}
	fmt.Printf("%s 你有恶意刷票行为,现在取消你的投票资格\n",user)
}
//拉入黑名单
type BlackVoteState struct {}
func (bvs *BlackVoteState)Vote(user string,voteItem string,voteManager *VoteManager)  {
	fmt.Printf("%s 由于你恶意刷票行为，现在已将你拉入黑名单，将禁止你登录使用本系统\n",user)
}

type VoteManager struct {
	State VoteSate //持有状态处理对象
	MapVote map[string]string //记录用户投票结果,map[用户名称]投票选项
	MapVoteCount map[string]int //记录用户投票次数,map[用户名称]投票次数
}
func (vm *VoteManager)Vote(user string,voteItem string)  {
	oldVoteCount,ok:=vm.MapVoteCount[user]
	if ok==false {
		oldVoteCount=0
	}
	oldVoteCount=oldVoteCount+1
	vm.MapVoteCount[user]=oldVoteCount

	if oldVoteCount==1 {
		//vm.State=&NormalVoteState{}
		//扩展新功能,给投票成功的用户额外10个积分做奖励
		vm.State=&NormalVoteState2{
			&NormalVoteState{},
		}
	}else if oldVoteCount >1 && oldVoteCount<5 {
		vm.State=&RepeatVoteState{}
	}else if oldVoteCount >=5 && oldVoteCount <8 {
		vm.State=&SpiteVoteState{}
	}else if oldVoteCount >= 8{
		vm.State=&BlackVoteState{}
	}
	vm.State.Vote(user,voteItem,vm)
}
func NewVoteManager() *VoteManager  {
	return &VoteManager{
		MapVote:make(map[string]string),
		MapVoteCount:make(map[string]int),
	}
}