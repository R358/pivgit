// Copyright (c) 2014, R358.org
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
// list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// * Neither the name of pivgit nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

//
// PivotalTracker is a Registered Trademark of http://pivotallabs.com
//

package pivapi5

type PivGitConfig struct {
	Projects []Project `json:"projects"`
	Token    string    `json:"token"`
}

type Project struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name'`
	Version  int64   `json:"version"`
	StoryIds []int64 `json:"story_ids"`
}

type Story struct {
	Id           int64   `json:"id"`
	ProjectId    int64   `json:"project_id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	StoryType    string  `json:"story_type"`
	CurrentState string  `json:"current_state"`
	Estimate     float64 `json:"estimate"`
	AcceptedAt   string  `json:"accepted_at"`
	Deadline     string  `json:"deadline"`
	RequestById  int64   `json:"requested_by_id"`
	OwnerIds     []int64 `json:"owner_ids"`
	LabelIds     []int64 `json:"label_ids"`
	TaskIds      []int64 `json:"task_ids"`
	FollowerIds  []int64 `json:"follower_ids"`
	CommentIds   []int64 `json:"comment_ids"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	externalId   string  `json:"external_id"`
	url          string  `json:"url"`
	kind         string  `json:"kind"`
}

// 1186736,
